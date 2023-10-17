package github

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandleRateLimitBodyResponseFormatter(body []byte, header http.Header) error {
	r := string(body)
	if strings.Contains(r, "API rate limit exceeded") {
		// get github x-limit-remaining
		remaining := header.Get("X-RateLimit-Remaining")
		// display seconds human readable
		ratelimitReset := header.Get("X-RateLimit-Reset")
		// convert ratelimitReset to int
		ratelimitResetInt, err := strconv.Atoi(ratelimitReset)
		if err != nil {
			return errors.New(fmt.Sprintf("could not convert ratelimitReset to int: %v", err))
		}

		// convert ratelimitReset to time
		ratelimitResetTime := time.Unix(int64(ratelimitResetInt), 0)
		// get ratelimitReset in seconds
		ratelimitResetSeconds := ratelimitResetTime.Sub(time.Now()).Seconds()

		return errors.New(
			fmt.Sprintf(
				"could not get latest release version: %v x-ratelimit-remaining %v x-ratelimit-reset %v rate-limit-expires-in %v",
				string(body),
				remaining,
				ratelimitReset,
				ratelimitResetSeconds,
			),
		)
	}

	return nil
}
