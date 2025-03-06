import {request} from '@elasticview/plugin-sdk'

export function HelloAction(data:any) {
  return request({
    url:   `/api/HelloWorld`,
    method: 'post',
    data
  })
}
