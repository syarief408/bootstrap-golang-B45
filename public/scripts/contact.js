function submitData() {
  let name = document.getElementById("inputProjectName").value;
  let email = document.getElementById("inputProjectEmail").value;
  let number = document.getElementById("inputProjectNumber").value;
  let subject = document.getElementById("inputProjectSubject").value;
  let message = document.getElementById("inputProjectMessage").value;

  if (name == "") {
    return alert("di isi donk");
  } else if (email == "") {
    return alert("di isi donk");
  } else if (number == "") {
    return alert("di isi donk");
  } else if (subject == "") {
    return alert("di isi donk");
  } else if (message == "") {
    return alert("di isi donk");
  }

  console.log(name);
  console.log(email);
  console.log(number);
  console.log(subject);
  console.log(message);

  let emailReceiver = "syarief408@gmail.com";

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Hallo nama saya ${name}, ${message}, silahkan kontak ke nomor ${number}`;
  a.click();

  let student = {
    name,
    email,
    number,
    subject,
    message,
  };

  console.log(student);
}
