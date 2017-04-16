package main

import "strconv"

// Loadable interface for objects requiring form loading
type Loadable interface {
	LoadValue(name string, value []string)
}

// LoadModel Loads values from form into struct instance
func LoadModel(loadable Loadable, data map[string][]string) {
	for key, value := range data {
		loadable.LoadValue(key, value)
	}
}

// Populate instConfig struct with form values
func (i *instaBotConfig) LoadValue(name string, value []string) {
	var err error
	switch name {
	case "username":
		i.login = value[0]
	case "password":
		i.password = value[0]
	case "likesPerDay":
		i.like_per_day, err = strconv.Atoi(value[0])
	case "commentsPerDay":
		i.comments_per_day, err = strconv.Atoi(value[0])
	case "hashtags":
		i.tag_list = value
	case "hashtagsBlacklist":
		i.tag_blacklist = value
	case "usersBlacklist":
		i.user_blacklist = value
	case "maxLikesPerHashtag":
		i.max_like_for_one_tag, err = strconv.Atoi(value[0])
	case "followsPerDay":
		i.follow_per_day, err = strconv.Atoi(value[0])
	case "followTime":
		i.follow_time, err = strconv.Atoi(value[0])
	case "unfollowPerDay":
		i.unfollow_per_day, err = strconv.Atoi(value[0])
	case "unfollowBreakMin":
		i.unfollow_break_min, err = strconv.Atoi(value[0])
	case "unfollowBreakMax":
		i.unfollow_break_max, err = strconv.Atoi(value[0])
	case "logMod":
		i.log_mod, _ = strconv.Atoi(value[0])
	case "proxy":
		i.proxy = value[0]
	case "unwantedUsernameList":
		i.unwanted_username_list = value
	case "unfollowWhitelist":
		i.unfollow_whitelist = value
	}
	check(err)
}
