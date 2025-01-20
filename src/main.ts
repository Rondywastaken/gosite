import './style.css'

const btn = document.getElementById("link-btn") as HTMLButtonElement
const link = document.getElementById("link") as HTMLInputElement

function getLinkValue() {
    btn.addEventListener("click", () => {
        if (link.value != "") {
            return link.value
        }
        else { alert("Input field is empty!") }
    })
}
