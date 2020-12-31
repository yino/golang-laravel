<style lang="less">
  @import './index.less';
</style>
<template>
  <div>
    <div class="content">
      <div>
        <Input v-model="keywords.name" placeholder="请输入姓名,支持模糊搜索" style="width: 300px"/>
        <Input v-model="keywords.mobile" placeholder="请输入电话号码" style="width: 300px"/>
        <Button type="success" @click="search">搜索</Button>
        <Button type="warning" @click="reset">重置</Button>
      </div>
      <div>
        <Table :columns="columns1" :data="memberList"></Table>
      </div>
      <div class="right">
        <Page :total="meta.total" :page-size="meta.size" show-elevator show-total @on-change="changePage"/>
      </div>
    </div>

    <Modal
      v-model="addMemberIntegralModal"
      :title="addMemberIntegralModalTitle"
      footer-hide>
      <member-add-integral @on-success-valid="handleSubmit"></member-add-integral>
      <p slot="footer"></p>
    </Modal>
  </div>


</template>
<script>
  import {memberAddIntegralRecord, memberDelete,memberList} from "@/api/data"
  import {formatDate} from "@/common/date.js"
  import memberAddIntegral from "_c/member-add-integral"

  export default {
    components: {
      memberAddIntegral,
      formatDate,
    },
    data() {
      return {
        columns1: [
          {
            title: '姓名',
            key: 'name'
          },
          {
            title: '手机号',
            key: 'mobile'
          },
          {
            title: '积分',
            key: 'integral'
          },
          {
            title: '创建时间',
            key: 'date'
          },
          {
            title: '操作',
            key: 'action',
            fixed: 'right',
            width: 200,
            render: (h, params) => {
              return h('div', [
                h('Button', {
                  props: {
                    type: 'text',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      this.jumpEdit(this.memberList[params.index]["id"])
                    }
                  }
                }, '编辑'),
                h('Button', {
                  props: {
                    type: 'text',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      this.showAddMemberIntegralModal(this.memberList[params.index]["id"], this.memberList[params.index]["name"])
                    }
                  },
                }, '增加流水'),
                h('Button', {
                  props: {
                    type: 'text',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      this.jumpRecord(this.memberList[params.index]["id"])
                    }
                  },
                }, '查看流水'),
                h('Button', {
                  props: {
                    type: 'text',
                    size: 'small'
                  },
                  on: {
                    click: () => {
                      this.delete(this.memberList[params.index]["id"])
                    }
                  },
                }, '删除')
              ]);
            }
          }
        ],
        memberList: [],
        meta: {
          page: 1,
          size: 10,
          total: 0,
        },
        keywords: {
          name: "",
          mobile: "",
        },
        addMemberIntegralModal: false,
        addMemberIntegralModalTitle: "123123",
        addMemberIntegralId: 0,
      }
    },

    // function
    methods: {
      getList() {
        let page = this.meta.page,
          limit = this.meta.size;

        // 获取 member list
        memberList(page, limit, this.keywords.name, this.keywords.mobile).then(res => {
          let data = res.data.list,
            memberList = [];
          this.meta.size = res.data.meta.size;
          this.meta.total = res.data.meta.total;
          data.forEach(function (item) {
            let date = new Date(item["created_at"] * 1000);
            memberList.push({
              name: item["name"],
              mobile: item["mobile"],
              integral: item["integral"],
              date: formatDate(date, 'yyyy年MM月dd日 hh:mm:ss'),
              id: item["id"],
            });
          });
          this.memberList = memberList;
        })
      },
      changePage(page) {
        this.meta.page = page;
        this.getList();
      },
      search() {
        if (this.keywords.name.length == 0 && this.keywords.mobile.length == 0) {
          return false;
        }
        this.getList();
      },
      reset() {
        this.keywords.name = "";
        this.keywords.mobile = "";
        this.getList();
      },

      jumpEdit(id) {
        this.$router.push({name: 'member_edit', params: {id: id}})
      },
      jumpRecord(id) {
        this.$router.push({name: 'show_member_record_list', params: {id: id}})
      },
      showAddMemberIntegralModal(id, name) {
        let modal = this.addMemberIntegralModal;
        this.addMemberIntegralModal = !modal
        this.addMemberIntegralModalTitle = name + "-增加流水";
        this.addMemberIntegralId = id;
      },

      handleSubmit({money, integral, product_id}) {
        memberAddIntegralRecord(this.addMemberIntegralId, product_id, money, integral).then(res => {
          if (res.code == 200){
            this.$Message.success("添加成功")
            this.addMemberIntegralModal = false;
            this.getList();
          }else{
            this.$Message.error(res.msg);
          }
        })
      },
      delete(id){
        this.$Modal.confirm({
            title: "警告",
            content: "是否确定删除",
            okText:"确定",
            cancelText:"取消",
            onOk: async ()=>{
              await memberDelete(id).then(res=>{
                if (res.code == 200){
                  this.$Message.success(res.msg)
                  this.getList();
                }else{
                  this.$Message.error(res.msg)
                }
              })
            }
        } )

      },
    },

    // 渲染完成后加载
    mounted() {
      this.getList();
    },

    // 也没面渲染时加载
    created() {
      // this.getList();
    }
  }

</script>
