package model

type AuthenticationInput struct {
  Username string `json:"username"`
  Password string `json:"password" binding:"required"`
  Email string `json:"email" binding:"required"`
}