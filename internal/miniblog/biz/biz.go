package biz

import (
	"github.com/yshujie/miniblog/internal/miniblog/biz/article"
	"github.com/yshujie/miniblog/internal/miniblog/biz/auth"
	"github.com/yshujie/miniblog/internal/miniblog/biz/blog"
	"github.com/yshujie/miniblog/internal/miniblog/biz/module"
	"github.com/yshujie/miniblog/internal/miniblog/biz/section"
	"github.com/yshujie/miniblog/internal/miniblog/biz/user"
	"github.com/yshujie/miniblog/internal/miniblog/store"
)

// IBiz 业务接口，定义了 Biz 层需要实现的方法
type IBiz interface {
	UserBiz() user.IUserBiz
	AuthBiz() auth.IAuthBiz
	ModuleBiz() module.IModuleBiz
	SectionBiz() section.ISectionBiz
	ArticleBiz() article.IArticleBiz
	BlogBiz() blog.IBlogBiz
}

// biz 业务实现
type biz struct {
	ds store.IStore
}

// 确保 biz 实现了 IBiz 接口
var _ IBiz = (*biz)(nil)

// NewBiz 简单工厂函数，创建 biz 实例
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users 返回用户业务实例
func (b *biz) UserBiz() user.IUserBiz {
	return user.New(b.ds)
}

// AuthBiz 返回认证业务实例
func (b *biz) AuthBiz() auth.IAuthBiz {
	return auth.NewAuthBiz(b.ds)
}

// ModuleBiz 返回模块业务实例
func (b *biz) ModuleBiz() module.IModuleBiz {
	return module.New(b.ds)
}

// SectionBiz 返回章节业务实例
func (b *biz) SectionBiz() section.ISectionBiz {
	return section.New(b.ds)
}

// ArticleBiz 返回文章业务实例
func (b *biz) ArticleBiz() article.IArticleBiz {
	return article.New(b.ds)
}

// BlogBiz 返回博客业务实例
func (b *biz) BlogBiz() blog.IBlogBiz {
	return blog.New(b.ds)
}
