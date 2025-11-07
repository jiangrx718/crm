import React from 'react';
import { Breadcrumb, Card, Tabs, Table } from 'antd';
import { Link } from 'react-router-dom';

type StatRow = { key: string; metric: string; value: number };

const OrderStatistics: React.FC = () => {
  const [activeTab, setActiveTab] = React.useState<string>('today');
  const data: StatRow[] = [
    { key: '1', metric: '订单数', value: 123 },
    { key: '2', metric: '支付订单数', value: 98 },
    { key: '3', metric: '退款订单数', value: 5 },
    { key: '4', metric: '成交金额（￥）', value: 12345 },
  ];

  return (
    <div>
      <Card>
        <Breadcrumb style={{ marginBottom: 20 }} items={[{ title: <Link to="/home">首页</Link> }, { title: '订单管理' }, { title: '订单统计' }]} />

        <Tabs
          items={[
            { key: 'today', label: '今日' },
            { key: 'week', label: '本周' },
            { key: 'month', label: '本月' },
          ]}
          activeKey={activeTab}
          onChange={setActiveTab}
          style={{ marginBottom: 12 }}
        />

        <Table
          columns={[
            { title: '指标', dataIndex: 'metric' },
            { title: '数值', dataIndex: 'value' },
          ]}
          dataSource={data}
          pagination={false}
          rowKey="key"
        />
      </Card>
    </div>
  );
};

export default OrderStatistics;