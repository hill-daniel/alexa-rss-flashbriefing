# Alexa codecentric Flash Briefing Skill 
Provides code for a small AWS Lambda function, which retrieves data from the codecentric RSS feed
and parses it to JSON. 
Code for the infrastructure can be found at [GitHub](https://www.github.com/hill-daniel/alexa-rss-flashbriefing-infrastructure).
Take a look at the blog post: [RSS Feed mit Alexa Flash Briefing ausliefern](https://blog.codecentric.de/2018/11/rss-feed-mit-alexa-flash-briefing-skill-ausliefern/).

## Usage
The deploy.sh script will test, build and deploy the skill. Feel free to adapt variables to your needs.
Requires [AWS CLI](https://aws.amazon.com/de/cli/).

## Environment Variables
The following env var is set by terraform upon creation:
* ATOM_FEED_URL - the URL to retrieve the RSS Feed from