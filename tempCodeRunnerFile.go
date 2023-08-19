func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("contact_handler_info_receiver")
	fmt.Fprintln(w, "You can contact us at contact@example.com")
}