import axios from "axios";
import {url} from "@/main";


export function getNotes() {
    const user = JSON.parse(localStorage.getItem("User"))
    const authToken = btoa(`${user.Login}:${user.Password}`)
    const authOpts = axios.create({
        headers: {
            Authorization : `Basic ${authToken}`
        }
    })
    return authOpts.get(url + "/api/v1/notes?user_id=" + user.ID)
}

export function newNote(data) {
    const user = JSON.parse(localStorage.getItem("User"))
    const authToken = btoa(`${user.Login}:${user.Password}`)
    const authOpts = axios.create({
        headers: {
            Authorization : `Basic ${authToken}`
        }
    })
    data.UserID = user.ID
    return authOpts.post(url + "/api/v1/notes", data)
}

export function patchNote(data) {
    const user = JSON.parse(localStorage.getItem("User"))
    const authToken = btoa(`${user.Login}:${user.Password}`)
    const authOpts = axios.create({
        headers: {
            Authorization : `Basic ${authToken}`
        }
    })
    data.UserID = user.ID
    return authOpts.patch(url + "/api/v1/notes", data)
}

export function deleteNote(opt) {
    const user = JSON.parse(localStorage.getItem("User"))
    const authToken = btoa(`${user.Login}:${user.Password}`)
    const authOpts = axios.create({
        headers: {
            Authorization : `Basic ${authToken}`
        }
    })
    return authOpts.delete(url + "/api/v1/notes", {data: opt})
}