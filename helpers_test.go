package sqltyping

import (
	"testing"
)

func TestSanitizeString(t *testing.T) {
	t.Log("Testing SanitizeString")
	{
		dt := "MU'AS"
		SanitizeString(&dt)

		t.Log("Data: ", dt)
	}
	t.Log("Testing SanitizeString")
	{
		dt := "{ \"Id\": \"200923434300\", \"UserId\":\"MU'AZ\"}"
		SanitizeString(&dt)

		t.Log("Data: ", dt)
	}
	t.Log("Testing SanitizeString")
	{
		dt := "{ \"Id\": \"200923434300\", \"UserId\":\"MU%%AZ\"}"
		SanitizeString(&dt)

		t.Log("Data: ", dt)
	}
	t.Log("Testing SanitizeString")
	{
		dt := "{ \"Id\": \"200923434300\", \"UserId\":\"MU?AZ\"}"
		SanitizeString(&dt)

		t.Log("Data: ", dt)
	}
}
