// Seleciona o formulário
const form = document.querySelector('#form-cadastro');

// Seleciona os campos de input
const inputNome = document.querySelector('#nome');
const inputSobrenome = document.querySelector('#sobrenome');
const inputEmail = document.querySelector('#email');
const inputSenha = document.querySelector('#senha');
const inputConfirmaSenha = document.querySelector('#confirma-senha');

// Seleciona o botão de submit
const btnSubmit = document.querySelector('#btn-submit');

// Seleciona os elementos de erro
const errorNome = document.querySelector('#error-nome');
const errorSobrenome = document.querySelector('#error-sobrenome');
const errorEmail = document.querySelector('#error-email');
const errorSenha = document.querySelector('#error-senha');
const errorConfirmaSenha = document.querySelector('#error-confirma-senha');

// Adiciona o evento de submit ao formulário
form.addEventListener('submit', function (event) {
  // Verifica se o nome foi preenchido
  if (inputNome.value.trim() === '') {
    errorNome.textContent = 'Por favor, preencha seu nome';
    inputNome.focus();
    event.preventDefault();
  } else {
    errorNome.textContent = '';
  }

  // Verifica se o sobrenome foi preenchido
  if (inputSobrenome.value.trim() === '') {
    errorSobrenome.textContent = 'Por favor, preencha seu sobrenome';
    inputSobrenome.focus();
    event.preventDefault();
  } else {
    errorSobrenome.textContent = '';
  }

  // Verifica se o e-mail foi preenchido e é válido
  const emailRegExp = /\S+@\S+\.\S+/;
  if (inputEmail.value.trim() === '') {
    errorEmail.textContent = 'Por favor, preencha seu e-mail';
    inputEmail.focus();
    event.preventDefault();
  } else if (!emailRegExp.test(inputEmail.value)) {
    errorEmail.textContent = 'Por favor, informe um e-mail válido';
    inputEmail.focus();
    event.preventDefault();
  } else {
    errorEmail.textContent = '';
  }

  // Verifica se a senha foi preenchida e tem pelo menos 6 caracteres
  if (inputSenha.value.trim() === '') {
    errorSenha.textContent = 'Por favor, preencha sua senha';
    inputSenha.focus();
    event.preventDefault();
  } else if (inputSenha.value.length < 6) {
    errorSenha.textContent = 'A senha deve ter pelo menos 6 caracteres';
    inputSenha.focus();
    event.preventDefault();
  } else {
    errorSenha.textContent = '';
  }

  // Verifica se a confirmação de senha foi preenchida e é igual à senha
  if (inputConfirmaSenha.value.trim() === '') {
    errorConfirmaSenha.textContent = 'Por favor, confirme sua senha';
    inputConfirmaSenha.focus();
    event.preventDefault();
  } else if (inputConfirmaSenha.value !== inputSenha.value) {
    errorConfirmaSenha.textContent = 'As senhas não coincidem';
    inputConfirmaSenha.focus();
    event.preventDefault();
  } else {
    errorConfirmaSenha.textContent = '';
  }

  function redirectToLoginPage() {
    window.location.href = "../cadastro.html";
}

});


  

