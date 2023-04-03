package com.luke;

public class CloneTest implements Cloneable {
    private double pi = 3.14;

    @Override
    protected Object clone() {
        try {
            return super.clone();
        } catch (CloneNotSupportedException e) {
            throw new RuntimeException(e);
        }
    }

    public static void main(String[] args) throws CloneNotSupportedException {
        CloneTest obj1 = new CloneTest();
        CloneTest obj2 = (CloneTest) obj1.clone();
        obj1.pi = 3.1415926;
        System.out.println(obj1.pi);
        System.out.println(obj2.pi);
    }
}
