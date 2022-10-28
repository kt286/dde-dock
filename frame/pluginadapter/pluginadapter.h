/*
 * Copyright (C) 2022 ~ 2022 Deepin Technology Co., Ltd.
 *
 * Author:     donghualin <donghualin@uniontech.com>
 *
 * Maintainer: donghualin <donghualin@uniontech.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
#ifndef PLUGINADAPTER_H
#define PLUGINADAPTER_H

#include "pluginsiteminterface.h"
#include "pluginsiteminterface_v20.h"

#include <QObject>

/** 适配器，当加载到v20插件的时候，通过该接口来转成v23接口的插件
 * @brief The PluginAdapter class
 */

class PluginAdapter : public QObject, public PluginsItemInterface
{
    Q_OBJECT

public:
    PluginAdapter(PluginsItemInterface_V20 *pluginInter);
    ~PluginAdapter();

    const QString pluginName() const override;
    const QString pluginDisplayName() const override;
    void init(PluginProxyInterface *proxyInter) override;
    QWidget *itemWidget(const QString &itemKey) override;

    QWidget *itemTipsWidget(const QString &itemKey) override;
    QWidget *itemPopupApplet(const QString &itemKey) override;
    const QString itemCommand(const QString &itemKey) override;
    const QString itemContextMenu(const QString &itemKey) override;
    void invokedMenuItem(const QString &itemKey, const QString &menuId, const bool checked) override;
    int itemSortKey(const QString &itemKey) override;
    void setSortKey(const QString &itemKey, const int order) override;
    bool itemAllowContainer(const QString &itemKey) override;
    bool itemIsInContainer(const QString &itemKey) override;
    void setItemIsInContainer(const QString &itemKey, const bool container) override;

    bool pluginIsAllowDisable() override;
    bool pluginIsDisable() override;
    void pluginStateSwitched() override;
    void displayModeChanged(const Dock::DisplayMode displayMode) override;
    void positionChanged(const Dock::Position position) override;
    void refreshIcon(const QString &itemKey) override;
    void pluginSettingsChanged() override;
    PluginType type() override;
    PluginSizePolicy pluginSizePolicy() const override;

    QIcon icon(const DockPart &dockPart) override;
    PluginStatus status() const override;
    QString description() const override;

    void setItemKey(const QString &itemKey);

private:
    PluginsItemInterface_V20 *m_pluginInter;
    QString m_itemKey;
};

#endif // PLUGINADAPTER_H
