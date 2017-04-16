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

bot = InstaBot(user_blacklist={},login='admin@example.com',password='password',like_per_day=22,comments_per_day=21,tag_list=['Dignissimos sit facere eligendi cumque atque in odit adipisicing aliqua In voluptatem est'],tag_blacklist=['Impedit labore ullam similique numquam vel odio autem obcaecati eum irure dolorem culpa autem'],user_blacklist={},max_like_for_one_tag=67,follow_per_day=12,follow_time=7,unfollow_per_day=0,unfollow_break_min=70,unfollow_break_max=12,log_mod=91,proxy='Labore iste laborum Ad fugiat nulla rerum aut mollitia eum nemo perferendis aut nihil facilis sunt ut harum aperiam et',unwanted_username_list=['lisyfus'],unfollow_whitelist=['Temporibus sed sunt omnis aliquid voluptatem Enim in proident reprehenderit dolor non deleniti'])while True:
    mode = 0
    if mode == 0:
        bot.new_auto_mod()