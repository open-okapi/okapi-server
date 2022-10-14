package org.okapi.server

import org.okapi.lib.okhttp.sub.OkHttpMeta

/**
 * author chen.kang
 */
class ServerApplication {

}

fun main() {
    val result = org.okapi.lib.okhttp.Client.request(OkHttpMeta())
    println(result.toString())
}