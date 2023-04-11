const form = document.querySelector('form');
const loginBtn = document.querySelector('#login-btn');
const successPopup = document.querySelector('#success-popup');

function validateEmail(email) {
    const re = /\S+@\S+\.\S+/;
    return re.test(email);
}

form.addEventListener('submit', function(event) {
    event.preventDefault();
    const emailInput = document.querySelector('#email');
    const passwordInput = document.querySelector('#password');
    const emailValue = emailInput.value.trim();

    if (!validateEmail(emailValue)) {
        alert('Por favor, insira um endereço de e-mail válido.');
        emailInput.focus();
        return false;
    }

    successPopup.classList.remove('hidden');
});

function redirectToLoginPage() {
    window.location.href = "../cadastro.html";
}
  