<template>
  <Form ref="memberAddIntegral" :model="formItem"  @keydown.enter.native="handleSubmit">

    <FormItem label="请选择剧本" >
      <Select v-model="formItem.select" @on-change="selectChange">
        <Option :value="i" v-for="(item, i) in productList">{{item["name"]}}</Option>
      </Select>
    </FormItem>
    <FormItem label="请输入本次获得积分">
      <Input v-model="formItem.integral" type="number" placeholder="请输入本次获得积分"></Input>
    </FormItem>
    <FormItem label="请输入金额">
      <Input v-model="formItem.money" type="number" placeholder="请输入本次消费金额"></Input>
    </FormItem>
    <FormItem>
      <Button @click="handleSubmit" type="primary" long>确定</Button>
    </FormItem>
  </Form>
</template>

<script>
  import {productAllList} from "@/api/data.js"
  export default {
    name: "memberAddIntegral",
    data() {
      return {
        formItem: {
          money: 0,
          integral: 0,
          select: -1,
          memberId:0,
          productId:0,
        },
        productList : [],
      }
    },
    methods: {
      handleSubmit () {
        if (this.formItem.productId == 0 ){
          this.$Message.error("请选择剧本")
          return false;
        }
        this.$emit('on-success-valid', {
          money: this.formItem.money,
          integral: this.formItem.integral,
          product_id: this.formItem.productId,
        })
      },
      productAllList(){
        productAllList().then(res => {
            console.log(res);
            if (res.code == 200){
              this.productList = res.data;
            }else{
              this.$Message.error(res.msg);
            }
        });
      },
      selectChange(){
        this.formItem.productId = this.productList[this.formItem.select]['id']
        this.formItem.integral = this.productList[this.formItem.select]['integral']
      }
    },

    mounted() {
      this.productAllList()
    }
  }
</script>

