package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupWebSocketRoutes(router fiber.Router) {
	router.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer c.Close()
		
		log.Println("WebSocket client connected")
		
		for {
			messageType, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}
			
			log.Printf("Received: %s", msg)
			
			// Echo back for now
			if err := c.WriteMessage(messageType, msg); err != nil {
				log.Println("WebSocket write error:", err)
				break
			}
		}
		
		log.Println("WebSocket client disconnected")
	}))
}
