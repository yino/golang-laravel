import axios from '@/libs/api.request'

export const getTableData = () => {
  return axios.request({
    url: 'get_table_data',
    method: 'get'
  })
}

export const getDragList = () => {
  return axios.request({
    url: 'get_drag_list',
    method: 'get'
  })
}

export const errorReq = () => {
  return axios.request({
    url: 'error_url',
    method: 'post'
  })
}

export const saveErrorLogger = info => {
  return axios.request({
    url: 'save_error_logger',
    data: info,
    method: 'post'
  })
}

export const uploadImg = formData => {
  return axios.request({
    url: 'image/upload',
    data: formData
  })
}

export const getOrgData = () => {
  return axios.request({
    url: 'get_org_data',
    method: 'get'
  })
}

export const getTreeSelectData = () => {
  return axios.request({
    url: 'get_tree_select_data',
    method: 'get'
  })
}

//------------------------------------

export const homeCount = () => {
  return axios.request({
    url: 'corp/index',
    method: 'get'
  })
}

//--------------------member-----------------------
export const memberList = (page, limit, name = "", mobile = "") => {
  return axios.request({
    url: 'corp/member/index',
    method: 'get',
    params: {page: page, size: limit, name: name, mobile: mobile},
  })
}
export const memberAdd = (name, mobile, integral = 0) => {
  return axios.request({
    url: 'corp/member/add',
    method: 'post',
    data: {
      name: name,
      mobile: mobile,
      integral: integral
    },
  })
}
export const memberEdit = (id, name, mobile, integral = 0) => {
  return axios.request({
    url: 'corp/member/edit',
    method: 'post',
    data: {
      name: name,
      mobile: mobile,
      integral: integral,
      id: id,
    },
  })
}

export const memberInfo = (id) => {
  return axios.request({
    url: 'corp/member/info',
    method: 'get',
    params: {
      id: id,
    },
  })
}

export const memberAddIntegralRecord = (memberId, productId, money, integral) => {
  return axios.request({
    url: 'corp/member/add_integral',
    method: 'post',
    data: {
      product_id: parseInt(productId),
      member_id: parseInt(memberId),
      money: parseInt(money),
      integral: parseInt(integral),
    },
  })
}

export const memberDelete = (id) => {
  return axios.request({
    url: 'corp/member/delete',
    method: 'get',
    params: {
      id: id,
    },
  })
}
//--------------------product-----------------------
export const productAllList = () => {
  return axios.request({
    url: 'corp/product/all_list',
    method: 'get',
  })
}

export const productList = (page, limit, name = "") => {
  return axios.request({
    url: 'corp/product/index',
    method: 'get',
    params: {page: page, size: limit, name: name},
  })
}
export const productAdd = (name, integral = 0, path="") => {
  return axios.request({
    url: 'corp/product/add',
    method: 'post',
    data: {
      name: name,
      integral: integral,
      path: path,
    },
  })
}
export const productEdit = (id, name, integral = 0,path="") => {
  return axios.request({
    url: 'corp/product/edit',
    method: 'post',
    data: {
      name: name,
      integral: integral,
      path: path,
      id: id,
    },
  })
}

export const productInfo = (id) => {
  return axios.request({
    url: 'corp/product/info',
    method: 'get',
    params: {
      id: id,
    },
  })
}
export const productDelete = (id) => {
  return axios.request({
    url: 'corp/product/delete',
    method: 'get',
    params: {
      id: id,
    },
  })
}

// ---------------------消费记录------------------
export const recordList = (page, limit,memberId=0) => {
  return axios.request({
    url: 'corp/record/index',
    method: 'get',
    params: {
      member_id: memberId,
      page: page,
      size: limit,
    },
  })
}
