import {request} from "../plugin_sdk/sdk";

export function DbInsert(data:any) {
    return request({
        url:   `/api/DbInsert`,
        method: 'post',
        data
    })
}

export function CleanAction(data:any) {
    return request({
        url:   `/api/DbDelete`,
        method: 'post',
        data
    })
}

export function ListAction(data:any) {
    return request({
        url:   `/api/DbSearch`,
        method: 'post',
        data
    })
}
