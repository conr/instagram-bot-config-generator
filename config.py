#!/usr/bin/env python
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

bot = InstaBot(login='lyladaqady',password='password',like_per_day=24,comments_per_day=1,tag_list=['Expedita sequi omnis delectus natus voluptatem porro officia quia nostrud amet soluta earum cum et sint neque magna non'],tag_blacklist=['Modi nisi sunt ducimus et dolore exercitation'],user_blacklist=['Laboriosam recusandae Voluptate velit omnis quam inventore quia nihil sapiente suscipit eveniet esse error'],max_like_for_one_tag=49,follow_per_day=23,follow_time=36,unfollow_per_day=0,unfollow_break_min=9,unfollow_break_max=19,log_mod=79,proxy='Repudiandae fugiat in autem error iste velit vitae praesentium sequi odit sunt dolore sint',unwanted_username_list=['hazagodij'],unfollow_whitelist=['Rerum qui consequatur sed qui nihil voluptatem perspiciatis ut'])while True:
    mode = 0
    if mode == 0:
        bot.new_auto_mod()