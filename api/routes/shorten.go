package  routes

import(
    "time"
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

    if !govalidator.IsURL(body.URL){
        return  c.Status(fiber.StatusBadRequest.JSON(fiber.Map("error": "Invalid URL")))
    }

    if !helpers.RemoveDomainError(body.URL){
        return c.Status(fiber.StatusBadRequest.JSON(fiber.Map("error": "Invalid Domain")))
    }

    body.URL = helpers.EnforceHTTP(body.URL)
}


