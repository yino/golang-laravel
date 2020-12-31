<template>
  <div>
    <Row :gutter="20">
      <i-col :xs="12" :md="8" :lg="6" v-for="(infor, i) in inforCardData" :key="`infor-${i}`"
             style="height: 120px;padding-bottom: 10px;">
        <infor-card shadow :color="infor.color" :icon="infor.icon" :icon-size="36">
          <count-to :end="infor.count" count-class="count-style"/>
          <p>{{ infor.title }}</p>
        </infor-card>
      </i-col>
    </Row>
    <Row :gutter="20" style="margin-top: 10px;">
      <!--      <i-col :md="24" :lg="16" style="margin-bottom: 20px;">-->
      <Card shadow>
        <div class="Echarts">
          <div id="echarts" style="height:300px;"></div>
        </div>
      </Card>
      <!--      </i-col>-->
    </Row>
  </div>
</template>

<script>
  import InforCard from '_c/info-card'
  import CountTo from '_c/count-to'
  import {ChartBar, ChartPie} from '_c/charts'
  import Example from './example.vue'
  import {homeCount} from '@/api/data'
  import echarts from "echarts";

  export default {
    name: 'home',
    components: {
      InforCard,
      CountTo,
      ChartPie,
      ChartBar,
      Example
    },
    data() {
      return {
        inforCardData: [],
      }
    },
    mounted() {
      this.tabCount();
    },
    methods: {

      myEcharts(xArr, yArr) {
        // 基于准备好的dom，初始化echarts实例
        let myChart = echarts.init(document.getElementById('echarts'));

        // 指定图表的配置项和数据
        let option = {
          title: {
            text: '最近7天开桌次数'
          },
          xAxis: {
            type: 'category',
            data: xArr
          },
          yAxis: {
            type: 'value'
          },
          series: [{
            data: yArr,
            type: 'bar'
          }]
        };
        // 使用刚指定的配置项和数据显示图表。
        myChart.setOption(option);
      },

      tabCount() {
        const _this = this;

        homeCount().then(res => {
          let data = res.data;
          _this.inforCardData = [
            {title: '总会员', icon: 'md-person-add', count: data["member_count"], color: '#2d8cf0'},
            {title: '剧本总数', icon: 'md-locate', count: data["product_count"], color: '#19be6b'},
            {title: '总开桌次数', icon: 'md-help-circle', count: data["count"], color: '#ff9900'},
            {title: '总金额', icon: 'md-share', count: data["total_money"] / 100, color: '#ed3f14'},
          ];

          let xArr = [],
            yArr = [],
            weekGroupCount = data["week_group_count"];
          weekGroupCount.forEach(item => {
            xArr.push(item["name"]);
            yArr.push(item["value"]);
          })
          _this.myEcharts(xArr, yArr);
        })
      },
    },
  }
</script>

<style lang="less">
  .count-style {
    font-size: 50px;
  }
</style>
