/*
 * Copyright (C) 2015 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     sbw <sbw@sbw.so>
 *
 * Maintainer: sbw <sbw@sbw.so>
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

/*
 * This file was generated by qdbusxml2cpp version 0.8
 * Command line was: qdbusxml2cpp -c DBusMenuManager -p dbusmenumanager com.deepin.menu.Manager.xml
 *
 * qdbusxml2cpp is Copyright (C) 2015 Digia Plc and/or its subsidiary(-ies).
 *
 * This is an auto-generated file.
 * Do not edit! All changes made to it will be lost.
 */

#ifndef DBUSMENUMANAGER_H_1436158928
#define DBUSMENUMANAGER_H_1436158928

#include <QtCore/QObject>
#include <QtCore/QByteArray>
#include <QtCore/QList>
#include <QtCore/QMap>
#include <QtCore/QString>
#include <QtCore/QStringList>
#include <QtCore/QVariant>
#include <QtDBus/QtDBus>

/*
 * Proxy class for interface com.deepin.menu.Manager
 */
class DBusMenuManager: public QDBusAbstractInterface
{
    Q_OBJECT
public:
    static inline const char *staticServerPath()
    { return "com.deepin.menu"; }
    static inline const char *staticInterfacePath()
    { return "/com/deepin/menu"; }
    static inline const char *staticInterfaceName()
    { return "com.deepin.menu.Manager"; }

public:
    explicit DBusMenuManager(QObject *parent = 0);

    ~DBusMenuManager();

public Q_SLOTS: // METHODS
    inline QDBusPendingReply<QDBusObjectPath> RegisterMenu()
    {
        QList<QVariant> argumentList;
        return asyncCallWithArgumentList(QStringLiteral("RegisterMenu"), argumentList);
    }

    inline QDBusPendingReply<> UnregisterMenu(const QString &menuObjectPath)
    {
        QList<QVariant> argumentList;
        argumentList << QVariant::fromValue(menuObjectPath);
        return asyncCallWithArgumentList(QStringLiteral("UnregisterMenu"), argumentList);
    }

Q_SIGNALS: // SIGNALS
};

namespace com {
  namespace deepin {
    namespace menu {
      typedef ::DBusMenuManager Manager;
    }
  }
}
#endif
