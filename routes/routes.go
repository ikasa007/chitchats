package routes

import (
	"github.com/tim/chitchat/handlers"
	"github.com/tim/chitchat/handlers/home"
	"github.com/tim/chitchat/handlers/posts"
	"github.com/tim/chitchat/handlers/threads"
	"net/http"
)

type WebRoute struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute


var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		home.Index,
	},
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	{
		"signup",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	{
		"threads",
		"GET",
		"/thread/read",
		threads.GetThreads,
	},
	{
		"threads",
		"GET",
		"/thread/new",
		threads.CreateThreads,
	},
	{
		"threads",
		"POST",
		"/thread/create",
		threads.PostsThreads,
	},
	{
		"postThread",
		"POST",
		"/thread/post",
		posts.PostThread,
	},
	{
		"error",
		"GET",
		"/err",
		home.Err,
	},
}

