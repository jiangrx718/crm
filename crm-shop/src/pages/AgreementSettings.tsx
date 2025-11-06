import React, { useState } from 'react';
import { Breadcrumb, Card, Tabs, Space, Button, Input } from 'antd';
import {
  BoldOutlined,
  ItalicOutlined,
  UnderlineOutlined,
  OrderedListOutlined,
  UnorderedListOutlined,
  LinkOutlined,
  AlignLeftOutlined,
  AlignCenterOutlined,
  AlignRightOutlined,
} from '@ant-design/icons';

const { TextArea } = Input;

const agreementTabs = [
  { key: 'vip', label: '付费会员协议' },
  { key: 'agent', label: '代理商协议' },
  { key: 'privacy', label: '隐私协议' },
  { key: 'user', label: '用户协议' },
  { key: 'law', label: '法律协议' },
  { key: 'point', label: '积分协议' },
  { key: 'distribution', label: '分销协议' },
];

const AgreementSettings: React.FC = () => {
  const [activeKey, setActiveKey] = useState(agreementTabs[0].key);
  const [content, setContent] = useState('');

  const save = () => {
    // 这里可对当前 activeKey 的协议内容进行保存
  };

  return (
    <div className="page-container">
      <Breadcrumb items={[{ title: '首页' }, { title: '系统设置' }, { title: '协议设置' }]} />

      <Card style={{ marginTop: 16 }}>
        <Tabs
          activeKey={activeKey}
          items={agreementTabs.map(t => ({ key: t.key, label: t.label }))}
          onChange={setActiveKey}
        />

        <Space align="center" style={{ marginBottom: 12 }}>
          <span style={{ color: '#999' }}>HTML</span>
          <BoldOutlined />
          <ItalicOutlined />
          <UnderlineOutlined />
          <OrderedListOutlined />
          <UnorderedListOutlined />
          <LinkOutlined />
          <AlignLeftOutlined />
          <AlignCenterOutlined />
          <AlignRightOutlined />
        </Space>

        <TextArea
          value={content}
          onChange={e => setContent(e.target.value)}
          placeholder="请在此输入协议内容..."
          style={{ height: 420 }}
        />

        <div style={{ marginTop: 12 }}>
          <Button type="primary" onClick={save}>保存</Button>
        </div>
      </Card>
    </div>
  );
};

export default AgreementSettings;