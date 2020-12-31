<style lang="less">
  @import './form.less';
</style>
<template>
  <Form :model="formItem" :label-width="80" :rules="ruleInline">
    <FormItem label="姓名">
      <Input v-model="formItem.name" placeholder="姓名"></Input>
    </FormItem>
    <FormItem label="手机号码">
      <Input v-model="formItem.mobile" placeholder="手机号码"></Input>
    </FormItem>
    <FormItem label="初始积分">
      <Input v-model="formItem.integral" placeholder="初始积分" type="number"></Input>
    </FormItem>
    <FormItem>
      <Button type="primary" @click="submit">提交</Button>
      <Button style="margin-left: 8px" @click="reset">取消</Button>
    </FormItem>
  </Form>
</template>

<script>
  import {memberAdd, memberInfo} from "@/api/data.js"
  import {memberEdit} from "../../../api/data";

  export default {
    data() {
      return {
        formItem: {
          name: '',
          mobile: '',
          integral: 0,
          id: 0,
        },
        ruleInline: {
          name: [
            {required: true, message: '请输入姓名', trigger: 'blur'}
          ],
          mobile: [
            {required: true, message: '请输入手机号', trigger: 'blur'}
          ],
          integral: [
            {required: true, message: '请输入积分', trigger: 'blur'},
            {type: 'number', message: '积分必须是数字', trigger: 'blur'},
          ],
        },
        type: "add",
        id: 0,
      }
    },

    methods: {
      submit() {
        let integral = parseInt(this.formItem.integral)
        // 参数验证
        if (this.formItem.name.length == 0 || this.formItem.name.length == " ") {
          this.$Message.error("请输入姓名");
          return false;
        } else if (this.formItem.mobile.length == 0 || this.formItem.mobile.length == " ") {
          this.$Message.error("请输入电话号码");
          return false;
        } else if (integral < 0) {
          this.$Message.error("积分不能小于0");
          return false;
        }

        // 提交
        if (this.type == "add"){
          memberAdd(this.formItem.name, this.formItem.mobile, integral).then(res => {
            if (res.code != 200) {
              this.$Message.error(res.msg);
              return false;
            } else {
              this.reset();
              this.$Message.success(res.msg);
            }
          });
        }else{
          memberEdit(this.formItem.id, this.formItem.name, this.formItem.mobile, integral).then(res => {
            if (res.code != 200) {
              this.$Message.error(res.msg);
              return false;
            } else {
              this.getInfo();
              this.$Message.success(res.msg);
            }
          });
        }

      },
      reset() {
        this.formItem = {
          input: '',
          mobile: '',
          integral: 0,
        }
      },
      getInfo() {
        if (this.type == "add") {
          return false;
        }
        let id = this.id
        memberInfo(id).then(res => {
          console.log( res.data);

          if (res.code == 200) {
            this.formItem = {
              name: res.data.name,
              mobile: res.data.mobile,
              integral: res.data.integral,
              id: res.data.id,
            }
          }
        });
      }
    },

    mounted() {
      // 验证是否类型
      var id = this.$route.params.id;
      if (id != undefined) {
        this.type = "edit"
        this.id = id;
        this.getInfo()
      } else {
        this.type = "add"
      }
    }
  }
</script>
