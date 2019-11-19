# ICT-DEV-DT-18 Starter
Codebase voor module ICT-DEV-DT-18

## Introductie

Dit is de basis van de applicatie waar je gedurende de module Development & Tooling mee aan de slag gaat.

Tijdens de module werk je met behulp van Scrum in sprints aan nieuwe versies van deze _Pizza applicatie_.

Op Github, onder het tabje _Issues_ vind je een lange lijst van bugs, fixes en nieuwe features. 

> **LET OP: De focus van deze module ligt voornamelijk op de manier van samenwerken in een ontwikkelteam met het gebruik van GIT. Daarnaast ga je doorontwikkelen op de basis van Go **

Voor de deadline in week 9 lever je je groepsdocument in op ItsLearning. Je laatste commit die is gedaan voor de deadline geldt als laatste, ingeleverde versie. Let er dus op dat je laatste commit goed werkt, gezien je geen nieuwe commit kunt doen na de deadline!

## Installatie

Om het huidige path toe te voegen aan je `$GOROOT` environment variable voer je de script `init.ps1` uit met PowerShell (rechtermuisknop).

Om het programma uit te voeren voer je het commando `go run .` uit in de `src` folder (`cd src`). Hierna is de webserver op poort `5000` beschikbaar. Om de website te bezoeken ga je naar `http://localhost:5000`.

## Issues

De issue's zijn ingedeeld onder de volgende labels:

### `bug`

Tijdens de **eerste sprint** van de module ga je aan de slag met bugs. Deze moet je eerst kunnen reproduceren. Vervolgens schrijf je hier een test voor en los je de bug op.

> [Lijst van bugs](https://github.com/che-ict/ICT-DEV-DT-18/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Abug)

### `improvement`

Improvements zijn kleine aanpassingen die je doet aan de applicatie. Tijdens **sprint 2** van de module ga je aan de slag met deze issues.

> [Lijst van improvements](https://github.com/che-ict/ICT-DEV-DT-18/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Aimprovement)

### `feature`

Onder het label `feature` vind je compleet nieuwe wensen. Hierbij wordt ook wat meer van jullie verwacht: je maakt bij de nieuwe feature eerst de bijbehorende ontwerpen. Zo weet je als team wat er te doen staat. Uiteraard ontwikkel je de features op aparte branches. Nieuwe features pak je op in **sprint 3**. 

> [Lijst van features](https://github.com/che-ict/ICT-DEV-DT-18/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3Afeature)

## Deployment

Jullie webapplicatie draait natuurlijk lokaal op je laptop. Tijdens deze module zal er ook een live versie van jullie applicatie beschikbaar staan op [AWS](https://aws.amazon.com/)! Zodra je een commit doet op de `master` branch wordt de recentste code automatisch [gedeployed](https://www.atlassian.com/continuous-delivery/continuous-deployment).

Deze server en deploymentstappen worden door de TOA (Bob) gefaciliteerd. Hier hoef je dus niks voor te doen. Let er wel op: zorg tijdens een oplevering (of voor de deadline) dat je laatste commit goed werkt, omdat dit dus je live versie wordt! 

## Vragen

Heb je inhoudelijke vragen over bepaalde issues? Mention (@bobmulder of @bjornpostema) in je issue!
