# Lambda Metadata Scraper

![Go](https://img.shields.io/badge/Go-1.17-blue.svg?logo=go&longCache=true&logoColor=white&colorB=88C0D0&style=flat-square&colorA=4c566a)
![Goquery](https://img.shields.io/badge/Goquery-1.8.0-blue.svg?logo=go&longCache=true&logoColor=white&colorB=88C0D0&style=flat-square&colorA=4c566a)
![AWS Lambda](https://img.shields.io/badge/AWS--Lambda-1.27.1-blue.svg?logo=go&longCache=true&logoColor=white&colorB=88C0D0&style=flat-square&colorA=4c566a)
![GitHub Last Commit](https://img.shields.io/github/last-commit/google/skia.svg?style=flat-square&colorA=4c566a&colorB=a3be8c&logo=GitHub)
[![GitHub Issues](https://img.shields.io/github/issues/toddbirchard/lambda-metadata-scraper.svg?style=flat-square&colorA=4c566a&colorB=ebcb8b&logo=GitHub)](https://github.com/toddbirchard/lambda-metadata-scraper/issues)
[![GitHub Stars](https://img.shields.io/github/stars/toddbirchard/lambda-metadata-scraper.svg?style=flat-square&colorB=ebcb8b&colorA=4c566a&logo=GitHub)](https://github.com/toddbirchard/lambda-metadata-scraper/stargazers)
[![GitHub Forks](https://img.shields.io/github/forks/toddbirchard/lambda-metadata-scraper.svg?style=flat-square&colorA=4c566a&colorB=ebcb8b&logo=GitHub)](https://github.com/toddbirchard/lambda-metadata-scraper/network)

Lambda function which returns basic metadata for a single URL provided as a `?url` querystring parameter. Intended to be deployed as an endpoint served with Netlify functions.

### Example Usage

* HTTP Request: `GET`
* Endpoint: `https://hackersandslackers.com/.netlify/functions/scrape`
* Parameters: `?url=https://toddbirchard.com`

```bash
$ curl https://hackersandslackers.com/.netlify/functions/scrape?url=https://toddbirchard.com
```

**Response:**

```json
{
   "Title": "Todd Birchard: Engineering, Product, Technology.",
   "Image": "https://storage.googleapis.com/toddbirchard-cdn/2019/08/cover.jpeg",
   "Description": "Giant reptile giving technology a good name. Occasional tangents of mass destruction. Made in Silicon Alley.",
   "Favicon": "/icons/icon-48x48.png"
}
```
