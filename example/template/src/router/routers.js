import Main from '@/components/main'

/**
 * iview-admin中meta除了原生参数外可配置的参数:
 * meta: {
 *  title: { String|Number|Function }
 *         显示在侧边栏、面包屑和标签栏的文字
 *         使用'{{ 多语言字段 }}'形式结合多语言使用，例子看多语言的路由配置;
 *         可以传入一个回调函数，参数是当前路由对象，例子看动态路由和带参路由
 *  hideInBread: (false) 设为true后此级路由将不会出现在面包屑中，示例看QQ群路由配置
 *  hideInMenu: (false) 设为true后在左侧菜单不会显示该页面选项
 *  notCache: (false) 设为true后页面在切换标签后不会缓存，如果需要缓存，无需设置这个字段，而且需要设置页面组件name属性和路由配置的name一致
 *  access: (null) 可访问该页面的权限数组，当前路由设置的权限会影响子路由
 *  icon: (-) 该页面在左侧菜单、面包屑和标签导航处显示的图标，如果是自定义图标，需要在图标名称前加下划线'_'
 *  beforeCloseName: (-) 设置该字段，则在关闭当前tab页时会去'@/router/before-close.js'里寻找该字段名对应的方法，作为关闭前的钩子函数
 * }
 */

export default [
  {
    path: '/login',
    name: 'login',
    meta: {
      title: 'Login - 登录',
      hideInMenu: true
    },
    component: () => import('@/view/login/login.vue')
  },
  {
    path: '/',
    name: '_home',
    redirect: '/home',
    component: Main,
    meta: {
      hideInMenu: true,
      notCache: true
    },
    children: [
      {
        path: '/home',
        name: 'home',
        meta: {
          hideInMenu: true,
          title: '首页',
          notCache: true,
          icon: 'md-home'
        },
        component: () => import('@/view/single-page/home')
      }
    ]
  },
  {
    path: '/message',
    name: 'message',
    component: Main,
    meta: {
      hideInBread: true,
      hideInMenu: true
    },
    children: [
      {
        path: 'message_page',
        name: 'message_page',
        meta: {
          icon: 'md-notifications',
          title: '消息中心'
        },
        component: () => import('@/view/single-page/message/index.vue')
      }
    ]
  },
  {
    path: '/member',
    name: 'member',
    meta: {
      icon: 'logo-buffer',
      title: '会员管理'
    },
    component: Main,
    children: [
      {
        path: 'member_list',
        name: 'member_list',
        meta: {
          icon: 'md-arrow-dropdown-circle',
          title: '会员管理'
        },
        component: () => import('@/view/member/list/index.vue')
      },
      {
        path: 'member/add',
        name: 'member_add',
        meta: {
          icon: 'md-add-circle',
          title: '添加会员'
        },
        component: () => import('@/view/member/form/add.vue')
      },
      {
        path: 'member/edit/:id',
        name: 'member_edit',
        meta: {
          hideInMenu: true,
          icon: 'md-add-circle',
          title: '编辑会员'
        },
        component: () => import('@/view/member/form/add.vue')
      },
    ]
  },
  // product
  {
    path: '/product',
    name: 'product',
    meta: {
      icon: 'logo-buffer',
      title: '剧本管理'
    },
    component: Main,
    children: [
      {
        path: 'product_list',
        name: 'product_list',
        meta: {
          icon: 'md-arrow-dropdown-circle',
          title: '剧本管理'
        },
        component: () => import('@/view/product/list/index.vue')
      },
      {
        path: 'product/add',
        name: 'product_add',
        meta: {
          icon: 'md-add-circle',
          title: '添加剧本'
        },
        component: () => import('@/view/product/form/add.vue')
      },
      {
        path: 'product/edit/:id',
        name: 'product_edit',
        meta: {
          hideInMenu: true,
          icon: 'md-add-circle',
          title: '编辑会员'
        },
        component: () => import('@/view/product/form/add.vue')
      },
    ]
  },
  {
    path: '/record',
    name: 'record',
    meta: {
      icon: 'logo-buffer',
      title: '消费记录'
    },
    component: Main,
    children: [
      {
        path: 'record_list',
        name: 'record_list',
        meta: {
          icon: 'md-arrow-dropdown-circle',
          title: '消费记录'
        },
        component: () => import('@/view/record/list/index.vue')
      },
      {
        path: 'show_member_record_list/:id',
        name: 'show_member_record_list',
        meta: {
          icon: 'md-arrow-dropdown-circle',
          title: '查看用户消费记录',
          hideInMenu: true
        },
        component: () => import('@/view/record/list/index.vue')
      },
    ]
  },
  {
    path: '/401',
    name: 'error_401',
    meta: {
      hideInMenu: true
    },
    component: () => import('@/view/error-page/401.vue')
  },
  {
    path: '/500',
    name: 'error_500',
    meta: {
      hideInMenu: true
    },
    component: () => import('@/view/error-page/500.vue')
  },
  {
    path: '*',
    name: 'error_404',
    meta: {
      hideInMenu: true
    },
    component: () => import('@/view/error-page/404.vue')
  }
]
