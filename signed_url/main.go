package main

func main() {
	//jsonKey, err := ioutil.ReadFile(serviceAccount)
	//if err != nil {
	//	return "", fmt.Errorf("cannot read the JSON key file, err: %v", err)
	//}
	//
	//conf, err := google.JWTConfigFromJSON(jsonKey)
	//if err != nil {
	//	return "", fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	//}
	//
	//opts := &storage.SignedURLOptions{
	//	Scheme:         storage.SigningSchemeV4,
	//	Method:         "GET",
	//	GoogleAccessID: conf.Email,
	//	PrivateKey:     conf.PrivateKey,
	//	Expires:        time.Now().Add(15 * time.Minute),
	//}
	//
	//u, err := storage.SignedURL(bucketName, objectName, opts)
	//if err != nil {
	//	return "", fmt.Errorf("Unable to generate a signed URL: %v", err)
	//}
	//
	//fmt.Fprintln(w, "Generated GET signed URL:")
	//fmt.Fprintf(w, "%q\n", u)
	//fmt.Fprintln(w, "You can use this URL with any user agent, for example:")
	//fmt.Fprintf(w, "curl %q\n", u)
	//// [END storage_generate_signed_url_v4]
	//return u, nil
}
