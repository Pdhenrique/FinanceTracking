package com.personal.financeTracking.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.client.reactive.ReactorClientHttpConnector;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.netty.http.client.HttpClient;
import reactor.netty.tcp.SslProvider;

import javax.net.ssl.*;
import java.io.FileInputStream;
import java.security.KeyStore;

@Configuration
public class C6ClientConfig {

    public SSLContext sslContext() throws Exception {
         String path = "/certs/";
         char[] password = "".toCharArray();

         KeyStore ks = KeyStore.getInstance("PKCS12");
         ks.load(new FileInputStream(path), password);

        KeyManagerFactory kmf = KeyManagerFactory.getInstance("SunX509");
        kmf.init(ks, password);

        TrustManagerFactory tmf = TrustManagerFactory.getInstance("SunX509");
        tmf.init(ks);

        SSLContext ctx = SSLContext.getInstance("TLS");
        ctx.init(kmf.getKeyManagers(), tmf.getTrustManagers(), null);
        return ctx;
    }

    @Bean
    public WebClient c6WebClient(SSLContext sslContext){
        HttpClient reactor = HttpClient.create()
                .secure(ssl -> ssl.sslContext((SslProvider.ProtocolSslContextSpec) sslContext));

        return WebClient.builder()
                .clientConnector(new ReactorClientHttpConnector(reactor))
                .baseUrl("https://apih.c6bank.com.br")
                .build();
    }

}
