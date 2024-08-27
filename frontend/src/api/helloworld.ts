import {request} from "../plugin_sdk/sdk";

export function HelloAction(data:any) {
  return request({
    url:   `/api/HelloWorld`,
    method: 'post',
    data
  })
}
