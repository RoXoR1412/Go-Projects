package routes

import (
	"os"
	"strconv"
	"time"
	"github.com/RoXoR1412/shorten-url-fiber-redis/helpers"
	"github.com/RoXoR1412/shorten-url-fiber-redis/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/go-redis/redis/v8"
	"github.com/asaskevich/govalidator"
)

type request struct {
	URL					string			`json:"url"`
	CustomShort	    	string			`json:"short"`
	Expiry				time.Duration	`json:"expiry"`
}

type response struct {
	URL					string			`json:"url"`
	CustomShort			string			`json:"short"`
	Expiry				time.Duration	`json:"expiry"`
	XRateRemaining		int				`json:"rate_limit"`
	XRateLimitReset		time.Duration		`json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {

	body := new(request)

	if err := c.BodyParser(&body); err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	//implement rate limiting
	r2 := database.CreateClient(1)
	defer r2.Close()
	val, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil{
		_ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		//val, _ = r2.Get(database.Ctx, c.IP()).Result()
		valInt, _ := strconv.Atoi(val)
		if valInt <=0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "Rate limit exceeded ",
				"rate_limit_reset": limit/time.Nanosecond/time.Minute,
				"Limit": os.Getenv("API_QUOTA"),
		})
	}
	}

	//check if the input is URL or not
	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL"})
	}

	//check for domain error
	if !helpers.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"Domain not allower"})
	}

	//enforce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	var id string
	if body.CustomShort == ""{
		id= uuid.New().String()[:6]
	} else {
		id= body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(database.Ctx, id).Result()
	if val != ""{
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL custom short is already in use",
		})
	}

	if body.Expiry ==0{
		body.Expiry=24
	}

	err = r.Set(database.Ctx, id, body.URL,body.Expiry*3600*time.Second).Err()
	if err!= nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":"Unable to connect to server",
		})
	}

	resp := response{
		URL:			body.URL,
		CustomShort:	"",
		Expiry:			body.Expiry,
		XRateRemaining:	10,
		XRateLimitReset:30,
	}

	r2.Decr(database.Ctx, c.IP())

	val, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)
	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.CustomShort = os.Getenv("DOMAIN")+"/"+id

	return c.Status(fiber.StatusOK).JSON(resp)
}