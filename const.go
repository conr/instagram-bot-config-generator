package main

type instaBotConfig struct {
	login                  string
	password               string
	like_per_day           int
	comments_per_day       int
	tag_list               []string
	tag_blacklist          []string
	user_blacklist         []string
	max_like_for_one_tag   int
	follow_per_day         int
	follow_time            int
	unfollow_per_day       int
	unfollow_break_min     int
	unfollow_break_max     int
	log_mod                int
	proxy                  string
	unwanted_username_list []string
	unfollow_whitelist     []string
}

// Top and bottom parts of exmaple.py python script
const initString = `#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import sys
import time

sys.path.append(os.path.join(sys.path[0], 'src'))

from check_status import check_status
from feed_scanner import feed_scanner
from follow_protocol import follow_protocol
from instabot import InstaBot
from unfollow_protocol import unfollow_protocol

bot = InstaBot(`

const endString = `while True:
    mode = 0
    if mode == 0:
        bot.new_auto_mod()`
