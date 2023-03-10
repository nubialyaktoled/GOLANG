package app

import ("github.com/urfave/cli"
"log"
"net"
"fmt")


func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação linha de comando"
	app.Usage = "Busca Ips"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
	{ 
		Name: "ip",
		Usage: "busca ips de endereços na internet",
		Flags: flags,
		Action: buscarIps,
	}, 
	{
		Name: "servidores",
		Usage: "Busca o nome dos servidores na internet",
		Flags: flags,
		Action: buscarServidores,
	},
}

	return app

}

func buscarIps(c *cli.Context){
  host := c.String("host")
  ips, erro := net.LookupIP(host)
  if erro != nil {
	log.Fatal(erro)
}

for _, ip := range ips {
	fmt.Println(ip)
}
}

func buscarServidores(c *cli.Context){
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}

