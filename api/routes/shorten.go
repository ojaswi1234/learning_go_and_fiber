package  routes

import(
    "time"
    "github.com/ojaswi1234/learning_go_and_fiber/database"
)




type request struct{
    URL  string `json:"url"`
    CustomShort string `json:"short"`
    Expiry time.Duration `json:"expiry"`
}


type response struct {
    URL string `json:"url"`
    CustomShort string `json:"short"`
    Expiry time.Duration `json:"expiry"`
    XRateRemaining int `json:"rate_limit"`
    XRateLimitReset time.duration `json:"rate_limit_reset"`
}



func shorten_url(c *fiber.Ctx) error {
    body := new(request)


    if err := c.BodyParser(&body); err != nil {
        return s.Status(fiber.StatusBadRequest).JSON(fiber.Map("error": "connot parse JSON"))
    }


    //Implementing Rate Limiting

    r2 := database .CreateClient(1)

    defer r2.Close()


    val,err := r2.Get(database.Ctx,c.IP()).Result()
    if err == redis.Nil {
        _ = r2.Set(database.Ctx, c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
    }

    


    // Validating if URL is right
    if !govalidator.IsURL(body.URL){
        return  c.Status(fiber.StatusBadRequest.JSON(fiber.Map("error": "Invalid URL")))
    }

    // checking domain error
    if !helpers.RemoveDomainError(body.URL){
        return c.Status(fiber.StatusBadRequest.JSON(fiber.Map("error": "Invalid Domain")))
    }


    //Enforcing/Adding HTTP 
    body.URL = helpers.EnforceHTTP(body.URL)
}


