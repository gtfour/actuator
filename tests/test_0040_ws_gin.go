r := gin.New()
r.GET("/ws", func(c *gin.Context) {
    handler := websocket.Handler(EchoServer)
    handler.ServeHTTP(c.Writer, c.Req)
})
r.Run(":8080")

###


