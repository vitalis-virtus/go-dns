package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/miekg/dns"
)

type Record struct {
	Domain string `json:"domain"`
	IP     string `json:"ip"`
}

var dnsRecords = map[string]string{
	"example.com.":  "192.0.2.1",
	"test.com.":     "203.0.113.42",
	"namanraj.tech": "192.0.2.123",
}

func main() {
	go startDNSServer()
	startWebServer()
}

func startWebServer() {
	app := fiber.New()

	app.Get("/records", func(c *fiber.Ctx) error {
		return c.JSON(dnsRecords)
	})

	app.Post("/records", func(c *fiber.Ctx) error {
		record := new(Record)

		if err := c.BodyParser(&record); err != nil {
			return c.Status(400).SendString("Bad Request")
		}

		if record.Domain[len(record.Domain)-1] != '.' {
			record.Domain += "."
		}

		dnsRecords[record.Domain] = record.IP
		return c.Status(201).SendString("Success")
	})

	log.Printf("Starting web server on :3000...")
	log.Fatal(app.Listen(":3000"))
}

func startDNSServer() {
	server := &dns.Server{
		Addr: "0.0.0.0:9090",
		Net:  "udp",
	}

	dns.HandleFunc(".", handleDNSRequest)

	log.Printf("Starting DNS server on port 9090...")

	log.Fatal(server.ListenAndServe())
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true
	found := false

	for _, q := range r.Question {
		if ip, exists := dnsRecords[q.Name]; exists {
			rr, _ := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
			m.Answer = append(m.Answer, rr)
			found = true
		}
	}

	if !found {
		m.Rcode = dns.RcodeNameError // NXDOMAIN
	}

	w.WriteMsg(m)
}
