# WeatherAssignment
•	Got an API key from OpenWeatherMap by signing up on their website.
•	Used Go's HTTP package to make a request to the OpenWeatherMap API, providing the necessary parameters.
•	Parse the API response and format it into JSON.
•	In this file The http.Get function is used to make the API request, and ioutil.ReadAll reads the response body.
•	json.Unmarshal is used to parse the JSON response into the WeatherData struct.
•	json.MarshalIndent is used to format the struct into a pretty-printed JSON string.
