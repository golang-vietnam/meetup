var _ = Resource("subtemplates", func() {
	DefaultMedia(SubtemplateMedia)
	BasePath("/api/subtemplate")

	Action("show", func() {
		Description("Get subtemplate")
		Routing(GET("/:subTemplateID"))
		Params(func() {
			Param("subTemplateID", Integer)
		})
		Response(OK, SubtemplateMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	// ...
})
