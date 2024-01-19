let hamburgerIsOpen = false

function openMenu() {
    let hamburgerContainer = document.getElementById("hamburger-container")
    if (!hamburgerIsOpen) {
        hamburgerContainer.style.display = "flex"
        // efek animasi
        hamburgerIsOpen = true;
    } else {
        hamburgerContainer.style.display = "none"
        hamburgerIsOpen = false;
    }
}