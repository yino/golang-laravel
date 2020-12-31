<style lang="less">
  @import './index.less';
</style>
<template>
  <div>
    <div class="content">
<!--      <div>-->
<!--        <Input v-model="keywords.name" placeholder="请输入名称,支持模糊搜索" style="width: 300px"/>-->
<!--        <Button type="success" @click="search">搜索</Button>-->
<!--        <Button type="warning" @click="reset">重置</Button>-->
<!--      </div>-->
      <div>
        <Table :columns="columns1" :data="memberList"></Table>
      </div>
      <div class="right">
        <Page :total="meta.total" :page-size="meta.size" show-elevator show-total @on-change="changePage"/>
      </div>
    </div>
  </div>


</template>
<script>
  import {recordList} from "@/api/data.js"
  import {formatDate} from "@/common/date.js"

  export default {
    components: {
      formatDate,
      memberId:0,
    },
    data() {
      return {
        columns1: [
          {
            title: '剧本名称',
            key: 'product_name'
          },
          {
            title: '用户名称',
            key: 'member_name'
          },
          {
            title: '获得积分',
            key: 'integral'
          },
          {
            title: '消费金额',
            key: 'money'
          },
          {
            title: '创建时间',
            key: 'date'
          },
        ],
        memberList: [],
        meta: {
          page: 1,
          size: 10,
          total: 0,
        },
        keywords: {
          name: "",
        },
      }
    },

    // function
    methods: {
      getList() {
        let page = this.meta.page,
          limit = this.meta.size;
        // 获取 member list
        recordList(page, limit, this.memberId).then(res => {
          let data = res.data.list,
            memberList = [];
          this.meta.size = res.data.meta.size;
          this.meta.total = res.data.meta.total;
          data.forEach(function (item) {
            let date = new Date(item["created_at"] * 1000);
            memberList.push({
              member_name: item["member_name"],
              mobile_mobile: item["mobile_mobile"],
              product_name: item["product_name"],
              money: item["money"],
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
        this.$router.push({name: 'product_edit', params: {id: id}})
      },


      delete(id){
        this.$Modal.confirm({
          title: "警告",
          content: "是否确定删除",
          okText:"确定",
          cancelText:"取消",
          onOk: async ()=>{
            await productDelete(id).then(res=>{
              if (res.code == 200){
                this.$Message.success(res.msg)
                this.getList();
              }else{
                this.$Message.error(res.msg)
              }
            })
          }
        } )
      }
    },

    // 渲染完成后加载
    mounted() {
      this.memberId = this.$route.params.id;
      this.getList();
    },

    // 也没面渲染时加载
    created() {
      // this.getList();
    }
  }

</script>
