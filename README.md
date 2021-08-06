<h1 align="center"> ChangeTower </h1>

![image](https://user-images.githubusercontent.com/86517035/128545450-d6eb6163-f698-43f2-af81-37b6a3a6dd5b.png)

[![MIT License](https://img.shields.io/github/license/Dc4ts/ChangeTower?color=blue)](https://github.com/Dc4ts/ChangeTower/blob/main/LICENSE)
![top language](https://img.shields.io/github/languages/top/dc4ts/ChangeTower?color=%23000000)
[![GitHub issues](https://img.shields.io/github/issues/dc4ts/ChangeTower)](https://github.com/Dc4ts/ChangeTower/issues)
</br>
</br>
ChangeTower is intended to help you watch changes in web pages and get notified of any changes written in Go. This tool is good to know the web pages are updated something or not to work on the new site before others

## Installation Instructions

ChangeTower requires go1.16+ to install successfully. Run the following command to get the repo
```sh
go get -v github.com/Dc4ts/ChangeTower
```
If you haven't got Go in your OS you can use the pre-build version by going to [release](https://github.com/Dc4ts/ChangeTower/releases)

## Usage/Help
<details>
<summary>👉 ChangeTower help menu 👈</summary><br>
<pre>
Usage of ./ChangeTower:
  -s (default false)
    silent mode (no banner if you want to parse final result to other tools)
  -u (default false)
    example of usage (how to use ChangeTower)
  -d (default false)
    dont log the data (logged data)
  -w (default false)
    without color (if your terminal cant print the colors)
  -l (default false)
    return links without any tag (dont effect in log file)
</pre>
</details>

## Example commands

+ Print banner without color if your OS can't print the ASCII colors.
```sh
cat links.txt | ./ChangeTower -w
```

+ Show the results in silent mode to parse the final result into another tools.
```sh
cat links.txt | ./ChangeTower -s
```

+ Run ChangeTower without logging the results.
```sh
echo "https://example.com" | ./ChangeTower -d
```

+ Show the results without any stuff.
```sh
cat links.txt | ./ChangeTower -s -l
```

## How to run automaticly (for linux users)

Here is a sample wich use cronjob for run every 2 hours and [notify by ProjectDiscovery](https://github.com/projectdiscovery/notify) to send result in our telegram bot
+ Make a bot in [botfather](https://t.me/BotFather) and get the API-KEY
+ Send a message to [getidsbot](https://t.me/getidsbot) and get your CHAT-ID
+ Get notify from [HERE](https://github.com/projectdiscovery/notify)
+ Add your target links in a txt file (E.x. links.txt)
+ Type this
```
crontab -e
```
+ In the end, you can send ChangeTower results on every social media. In this example we send our results with the telegram bot. (Edit the command below and add your API-KEY, CHAT-ID and etc)
```cron
* */2 * * * cat ~/links.txt | ChangeTower -s | notify -telegram -telegram-api-key 18620*****:AA*******tu-7qDyQUy7VcQs******5ZAPE -telegram-chat-id 1234567890
```
+ Now every 2 hours its give to ChangeTower and run ChangeTower without a banner to don't send it by notify and notify will send all `cat ~/links.txt | ChangeTower -s` result in the telegram bot to the CHAT-ID is given.

## Contributing

Contributions are always welcome!
See [contributing.md](https://github.com/Dc4ts/ChangeTower/blob/main/contributing.md) for ways to get started.
Please adhere to this project's [code of conduct](https://github.com/Dc4ts/ChangeTower/blob/main/CODE_OF_CONDUCT.md).

## License

ChangeTower is made with ♥️ by [Osb0rn3](https://t.me/BotFather) & [Alins_ir](https://t.me/alins_ir) and it is released under the [MIT license](https://choosealicense.com/licenses/mit/)
