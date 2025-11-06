import React, { useEffect, useRef, useState } from 'react';
import { Breadcrumb, Card, Tabs, Button } from 'antd';
import Quill from 'quill';
import 'quill/dist/quill.snow.css';

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
  const quillContainerRef = useRef<HTMLDivElement | null>(null);
  const quillRef = useRef<Quill | null>(null);

  // Quill 工具栏模块的简易类型，用于解决 addHandler 的类型报错
  type QuillToolbarModule = { addHandler: (name: string, handler: () => void) => void };

  const save = () => {
    // 这里可对当前 activeKey 的协议内容进行保存
  };

  // 初始化一次，避免重复创建导致出现两行工具栏
  useEffect(() => {
    if (!quillContainerRef.current) return;

    // 清理可能残留的 Quill DOM（工具栏/内容），避免重复渲染产生两行工具栏
    const wrapper = quillContainerRef.current.parentElement;
    wrapper?.querySelectorAll('.ql-toolbar').forEach(el => el.remove());
    quillContainerRef.current.innerHTML = '';

    quillRef.current = new Quill(quillContainerRef.current, {
      theme: 'snow',
      modules: {
        // 扩展常用控件：标题(H1/H2...)、字体大小、颜色/背景、删除线、引用、上下标、缩进、对齐、清除格式等
        toolbar: [
          [{ header: [1, 2, 3, 4, 5, 6, false] }, { font: [] }],
          [{ size: ['small', false, 'large', 'huge'] }],
          ['bold', 'italic', 'underline', 'strike'],
          [{ color: [] }, { background: [] }],
          [{ script: 'sub' }, { script: 'super' }],
          ['blockquote', 'code-block'],
          [{ list: 'ordered' }, { list: 'bullet' }, { indent: '-1' }, { indent: '+1' }],
          [{ align: '' }, { align: 'center' }, { align: 'right' }, { align: 'justify' }],
          ['link', 'image'],
          ['clean'],
        ],
      },
    });

    // 设置初始内容
    quillRef.current.root.innerHTML = content || '';

    // 图片上传（转Base64）
    const toolbar = quillRef.current.getModule('toolbar') as QuillToolbarModule;
    if (toolbar && typeof (toolbar as any).addHandler === 'function') {
      toolbar.addHandler('image', () => {
        const input = document.createElement('input');
        input.type = 'file';
        input.accept = 'image/*';
        input.onchange = async () => {
          const file = input.files?.[0];
          if (!file) return;
          const reader = new FileReader();
          reader.onload = () => {
            const range = quillRef.current!.getSelection(true);
            quillRef.current!.insertEmbed(range ? range.index : 0, 'image', reader.result as string, 'user');
          };
          reader.readAsDataURL(file);
        };
        input.click();
      });
    }

    // 内容变更监听
    quillRef.current.on('text-change', () => {
      setContent(quillRef.current!.root.innerHTML);
    });

    return () => {
      if (quillRef.current) {
        quillRef.current.off('text-change');
      }
    };
  }, []);

  return (
    <div className="page-container">
      <Breadcrumb items={[{ title: '首页' }, { title: '系统设置' }, { title: '协议设置' }]} />

      <Card style={{ marginTop: 16 }}>
        <Tabs
          activeKey={activeKey}
          items={agreementTabs.map(t => ({ key: t.key, label: t.label }))}
          onChange={setActiveKey}
        />

        <div style={{ border: '1px solid #e5e6eb', borderRadius: 6, overflow: 'hidden' }}>
          <div ref={quillContainerRef} style={{ height: 420 }} />
        </div>

        <div style={{ marginTop: 12 }}>
          <Button type="primary" onClick={save}>保存</Button>
        </div>
      </Card>
    </div>
  );
};

export default AgreementSettings;