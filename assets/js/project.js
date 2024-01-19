

// Fungsi untuk menentukan berapa hari, bulan, tahun, atau abad dari start_date dan end_date
// Parameter : start_date, end_date : string(format: yyyy-mm-dd)
const getDuration = (start_date, end_date) => {
    let start = new Date(start_date)
    let end = new Date(end_date)

    let duration = end - start

    // Jika durasi minus, maka return false
    if (duration < 0) {
        console.log(Math.floor((duration) / (1000 * 60 * 60 * 24)))
        return false
    }
    return true
}


// Prosedur untuk menghapus peringatan inputan kosong
// Parameter : id : string
const removeAlertEmpty = (id) => {
    document.getElementById(id).innerHTML = ''
}

// Prosedur untuk mengecek apakah form valid
// Parameter : event : event
const checkValidForm = (event) => {
    event.preventDefault()

    let start_date = document.getElementById('start-date').value
    let end_date = document.getElementById('end-date').value

    let isDataComplete = true

    duration = getDuration(start_date, end_date)
    if (duration == false) {
        document.getElementById('dateHelp').innerHTML = 'The end date cannot be earlier than the start date'
        isDataComplete = false
    } else {
        document.getElementById('dateHelp').innerHTML = ''
    }

    let iconTechCount = false
    if (document.getElementById('node-js').checked) {
        iconTechCount = true
    }
    else if (document.getElementById('react-js').checked) {
        iconTechCount = true
    }
    else if (document.getElementById('java').checked) {
        iconTechCount = true
    }
    else if (document.getElementById('go').checked) {
        iconTechCount = true
    }

    //  Cek apakah iconTechCount kosong atau tidak
    if (!iconTechCount) {
        document.getElementById('techHelp').innerHTML = 'Choose at least 1 technology'
        isDataComplete = false
    } else {
        document.getElementById('techHelp').innerHTML = ''
    }

    return isDataComplete;
}


// Fungsi untuk menambahkan project
// Parameter : event : event
const addProject = (event) => {
    if (!checkValidForm(event)) {
        return
    }
    
    const form = document.getElementById("form-project")
    const formData = new FormData(form)

    sendDataToServer(formData, "/", "POST")
}   

// Fungsi untuk mengubah project
// Parameter : event : event
const editProject = (event) => {
    if (!checkValidForm(event)) {
        return
    }

    const form = document.getElementById("form-project")
    const formData = new FormData(form)

    id = document.getElementById('id').value

    sendDataToServer(formData, "/edit-project/" + id, "POST")
}

// Fungsi untuk send data ke server
// Parameter : formData : FormData, url : string, method : string
function sendDataToServer(formData, url, method) {

    fetch(url,{
        action: url,
        method: method,
        body: formData  
    })
        .then(data => {
            console.log("breakpoint 1")
            console.log(data)
            // Pindahkan ke halaman project
            // window.location.href = "/"
            console.log("breakpoint 2")
        })
        .catch(error => {
            return console.log(error)
        })
}

console.log("Hello World")
console.log("Id =", document.getElementById('id').value)