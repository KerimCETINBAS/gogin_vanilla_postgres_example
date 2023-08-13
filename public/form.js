

document.addEventListener("DOMContentLoaded", function(){
    const loginForm     = document.querySelector("#loginForm")
    const messageForm   = document.querySelector("#messageForm")
    const registerForm  = document.querySelector("#registerForm")
    if (loginForm) {
            loginForm.addEventListener("submit", async (e)=> {
                e.preventDefault()
                const form = new FormData(e.target)
                const data = {
                    name: form.get("name"),
                    password: form.get("password")
                }
                const res = await fetch("http://localhost:8081/api/v1/auth/signin", {
                    method: "POST",
                    mode:"same-origin",
                    credentials: "same-origin",
                    body: JSON.stringify(data),
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json, text/html"
                    }
                }).then(data => {
                    if (data.status == 200) {
                        window.location.assign("/")
                    }
                })
        
            })
    }

    if(messageForm) {
        messageForm.addEventListener("submit", async (e) => {
            e.preventDefault()
            const message = new FormData(messageForm).get("message")
            
            const res = await fetch("/api/v1/messages", {
                method: "post",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    message
                })
            })
            console.log(res.status)
            if (res.status == 200) {
                window.location.reload()
            }

            
        })
    }
    

    if(registerForm) {
        registerForm.addEventListener("submit", async (e)=> { 
            e.preventDefault()
            const data = Object.fromEntries(new FormData(e.target))

            const res = await fetch("/api/v1/auth/signup", {
                method: "post",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data)
            })

            if(res.status == 200) {
                window.location.replace("/login")
            }
            
        })
    }
})