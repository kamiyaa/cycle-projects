FROM golang:latest

ARG SLACKBOT_HOME=/var/slack_bot
ARG WORK_DIR=/home/builds
ARG BOT_CONFIG=$SLACKBOT_HOME/bot.toml

ENV SLACKBOT_CONFIG $BOT_CONFIG

VOLUME $SLACKBOT_HOME

COPY / $WORK_DIR

RUN cd $WORK_DIR \
	&& go install .

ENTRYPOINT [ "./bin/slack-bot" ]
