import fetch from 'isomorphic-unfetch'

const baseURL = "http://localhost:8989"
// const uri = "https://www.ust.hk"

function fetchData(method, url, data) {
    return fetch(baseURL + url, {
        method: method,
        headers: {
            // 'Content-Type': 'application/json'
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: JSON.stringify(data),
    })
        .then(async r => {
            let result
            try {
                result = JSON.parse(await r.text())
            }
            catch (err) {
                throw "Unknown server response: " + r + "\nof error: " + err
            }

            if (!result.Success) {
                throw `Server error: ${r.status} ${r.statusText} - ${result.Error || ""}`
            }

            return result.Data
        })
        .then(r => {
            console.log(method, url, data, r)
            return r
        })
        .catch(err => {
            console.warn(method, url, data, err)
            throw err
        })
}

export default class API {

    static getBookList() {
        return fetchData('GET', "/book")
    }
    static getBook(id) {
        return fetchData('GET', "/book/" + id)
    }
    static getNote(id) {
        return fetchData('GET', "/note/" + id)
    }
}