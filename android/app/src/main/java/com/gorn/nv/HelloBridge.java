package com.gorn.nv;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;

import com.facebook.react.bridge.Callback;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;

import hello.Hello;

public class HelloBridge extends ReactContextBaseJavaModule {

    public HelloBridge (@Nullable ReactApplicationContext reactContext){
        super(reactContext);
    }
    @Override
    public String getName() {
        return "HelloBridge";
    }

    @ReactMethod
    public void sayHello (String name, Callback cb) {


        try {
            String hello = "Hello " + name;
            cb.invoke(null, hello);
            Hello.startListening();

        } catch (Exception err) {
            cb.invoke(err, null);
        }
    }
}
