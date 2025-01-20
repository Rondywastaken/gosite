import './style.css'

async function postLink(linkValue: string) {
    try {
        const response = await fetch('http://localhost:8080', {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ link: linkValue })
        });

        if (!response.ok) {
            throw new Error(`HTTP Error! status: ${response.status}`)
        }

        const data = await response.json()
        console.log("Success: ", data)
    } catch (error) {
        console.error("Error: ", error) 
    }
}

const btn = document.getElementById("link-btn") as HTMLButtonElement
const link = document.getElementById("link") as HTMLInputElement

btn.addEventListener("click", () => {
    if (link.value.trim() !== "") {
        postLink(link.value.trim())
    }
    else { alert("Input field is empty!") }
})
