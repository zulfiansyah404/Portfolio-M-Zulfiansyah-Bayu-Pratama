// Fungsi untuk mengecek apakah email valid atau tidak menggunakan regex
function validateEmail(email) {
    let re = /\S+@\S+\.\S+/;
    return re.test(email);
}

// Fungsi untuk mengirimkan data ke email
function submitData(data) {
    data.preventDefault()

    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phone = document.getElementById("phone").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value

    let objectData = {
        name: name,
        email,
        phone,
        subject,
        message
    }

    let arrayData = [name, email, phone, subject, message]

    console.log(objectData)

    if (name === "") {
        return alert('Nama harus diisi!')
    } else if (email === "") {
        return alert('Email harus diisi!')
    } else if (!validateEmail(email)) {
        return alert('Email tidak valid!')
    } else if (phone === "") {
        return alert('Phone harus diisi!')
    } else if (subject === "") {
        return alert('Subject harus diisi!')
    } else if (message === "") {
        return alert('Message harus diisi!')
    }

    const emailReceiver = "13521028@std.stei.itb.ac.id"

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo nama saya ${name},\n${message}, silahkan kontak saya di nomor berikut : ${phone}`
    a.click()
}

