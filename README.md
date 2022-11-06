## Drone Navigation System

## Additional Questions

1. For instrumentation, I will monitor the requests/per minute, the <br />latency of the requsts and also watch out for any CPU spike.<br /> My tool of choice here will be prometheus for exporting metrics and <br />grafana for visualization.

2. Throttling will be useful because you want to make sure that there <br />is a cool down period between the requests that the drone sends, <br />and a single drone doesn't end up hugging the entire resource. <br />Hence it will make sense, based on the expectation to limit the <br />number of requests a drone can send at any given period, eg not <br />more that 50 request per second.

3. To service several sectors at a time, the sectors ID can be kept in <br />a map and the `maths` service can take in this sector ID to  <br />perform it's calculation on the fly. Hence instead of passing the <br />sectorID when a new service is created, it can be passed across <br />directly to the `Calculate` function. to now make it look like:<br />
```go
func (ms *MathService) Calculate(ctx context.Context, sectorID float64, req *Request) float64
```
or the `Request` entity can be made to include the sectorID:
```go
type Request struct {
	CoordX   string `json:"x"`
	CoordY   string `json:"y"`
	CoordZ   string `json:"z"`
	Velocity string `json:"vel"`
}
```
and this will be populated before passing it on to the maths service.<br />thus making the calculation to be like:
```go
req.CoordX*req.SectorID + req.CoordY* req.sectorID + req.CoordZ*req.sectorID + req.Velocity
```

4. Two ways to accomplish this. 
a. The response package will have a new struct called `MomCorpResponse` where the server determines the origin of the request or the type of response, if it's `basic/mom`. The either `responses.Basic` is returned or `response.MomCorp` will be returned.
b. Instead of returning a response struct, a map is used. This way, we can dynamically set the key for the map and retain the value of the calculation. eg:
```go
response:=make(map[string]float64)
result:=ms.Calculate(ctx, req) // result of the calculation
switch (responseType){
    case mc: // mom-corp response
        response["location"] = result
        return
    default: // drone response
        response["loc"] = result
}
b, err:=json.Marshal(result) //we proceed to return result.
``` 

5. Versioning is the solution here. We can have `v1` and `v2` and either maintain the same calculation package or have a different <br />caluclation package. The handler can then pick whichever package it wants based on the version. <br />
There is already an interface `IMathService` and as long as the new <br /> calculation package implements this, all should be fine.  <br />
We can also have a map of the different versions, eg:
```go
version:=map[string]IMathService{
    "v1":maths.V1Service{},
    "v2":maths.V2Service{},
}

switch(req.Version) { //determine the version from the http-request
case "v1":
    return version["v1"].Calculate(ctx, req)
case "v2":
return version["v2"].Calculate(ctx, req)
default:
    return nil,"version not defined"
}
```

6. By having separate controlled environment to test those changes before customers can even have any access. <br />
Allowing for A/B Tests to also help with user feedback during release cycles. <br /> 
The deployment for example can be triggered by a CI/CD option but releases should be more intentional. Hence the rolling-out can be phased, either via a blue-green release or canary release.