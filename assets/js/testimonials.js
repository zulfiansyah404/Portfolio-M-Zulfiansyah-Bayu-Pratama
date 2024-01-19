const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.open("GET", "https://api.npoint.io/0443d18004bb2c4f8fdb", true)
    xhr.onload = function () {
        if (xhr.status === 200) {
            resolve(JSON.parse(xhr.responseText))
        } else if (xhr.status >= 400) {
            reject("Error loading data")
        }
    }
    xhr.onerror = function () {
        reject("Network error")
    }
    xhr.send()
})

let testiData = []

async function getData(rating) {
    try {
        const response = await promise
        console.log(response)
        testiData = response
        testimonialsAll()
    } catch (err) {
        console.log(err)
    }
}

getData()

// Prosedur pemanggilan semua data testimoni
function testimonialsAll() {
    let testimonialsHTML = ""

    testiData.forEach((items) => { 
        let icon = ""
        if (items.isCompany) {
            icon = `<i class="fas fa-building"></i>`
        } else {
            icon = `<i class="fas fa-user"></i>`
        }

        testimonialsHTML += `
        <div class="testimonial bg-dark">
            <div class="logo-test">
                ${icon}   
            </div>
            <div class="content-test">
                <img src="${items.image}" alt="">
                <div class="rating"> 
            `
        
        for (let i = 0; i < items.rating; i++) {
            testimonialsHTML += `
                <i class="fa-solid fa-star"></i>
            `
        }

        testimonialsHTML += `
                </div>
                <div class="comment">
                    <p class="text-white-50">
                        "${items.text}"
                    </p>
                </div>
                <h2 class="glacial-indifference">
                    - ${items.autors}
                </h2>
                
            </div>
        
        </div>
        `
    })     
    document.getElementById("testimonials").innerHTML = testimonialsHTML
}

testimonialsAll()

function checkSelectRating(rate) {
    if (rate.value == "all") {
        testimonialsAll()
    }
    else {
        FilterTestimonial(parseInt(rate.value))
    }
}

// Prosedur untuk memfilter data testimoni berdasarkan rating
function FilterTestimonial(rate) {
    console.log(rate)
    let filteredTestimonialHTML = ""

    const filteredData = testiData.filter((items) => {
        return items.rating === rate
    })

    filteredData.forEach((items) => { 
        let icon = ""
        if (items.isCompany) {
            icon = `<i class="fas fa-building"></i>`
        } else {
            icon = `<i class="fas fa-user"></i>`
        }

        filteredTestimonialHTML += `
        <div class="testimonial bg-dark">
            <div class="logo-test">
                ${icon}   
            </div>
            <div class="content-test">
                <img src="${items.image}" alt="">
                <div class="rating"> 
            `
        
        for (let i = 0; i < items.rating; i++) {
            filteredTestimonialHTML += `
                <i class="fa-solid fa-star"></i>
            `
        }

        filteredTestimonialHTML += `
                </div>
                <div class="comment">
                    <p class="text-white-50">
                        "${items.text}"
                    </p>
                </div>
                <h2 class="glacial-indifference">
                    - ${items.autors}
                </h2>
                
            </div>
        
        </div>
        `
    })     
    document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}