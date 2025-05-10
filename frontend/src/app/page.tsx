'use client'

import { motion } from "framer-motion";
import { useRouter } from "next/navigation";
import Head from "next/head";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";
import HomeAnimation from "@/components/HomeAnimation";
import { Group, Lock, MessageCircle, Smartphone } from "lucide-react";

export default function Home() {
  const router = useRouter();

  const features = [
    {
      title: "Real-Time Messaging",
      description: "Chat instantly with friends and colleagues, no delays.",
      icon: <MessageCircle size={40} />,
    },
    {
      title: "Group Chats",
      description: "Create group conversations for teams or social circles.",
      icon: <Group size={40} />,
    },
    {
      title: "Secure & Private",
      description: "End-to-end encryption keeps your chats safe.",
      icon: <Lock size={40} />,
    },
    {
      title: "Cross-Platform",
      description: "Access your chats on web, mobile, or desktop.",
      icon: <Smartphone size={40} />,
    },
  ];

  return (
    <>
      <Head>
        <title>ChatSphere</title>
        <meta name="description" content="Connect instantly with our secure chat platform" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="min-h-screen bg-black text-white flex flex-col items-center justify-center p-4 relative overflow-hidden">
        <HomeAnimation />
        <motion.div
          className="text-center relative z-10"
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
        >
          <h1 className="text-4xl md:text-6xl font-bold mb-4">
            Connect Instantly.
            <br />
            Chat Without Limits!
            <br />
            Secure, Fast, Seamless. 💬
          </h1>
          <p className="text-gray-400 max-w-xl mx-auto text-lg">
            Stay connected with friends, teams, or communities through our powerful chat platform.
          </p>
          <div className="mt-6 flex justify-center gap-4">
            <button
              className="bg-white text-black font-semibold px-6 py-3 rounded-lg shadow-lg hover:bg-gray-200 transition duration-200"
              onClick={() => router.push("/login")}
            >
              Start Chatting
            </button>
            <button
              className="border border-white bg-transparent text-white font-semibold px-6 py-3 rounded-lg shadow-lg hover:bg-gray-800 hover:text-white transition duration-200"
              onClick={() => router.push("/features")}
            >
              Explore Features
            </button>
          </div>
        </motion.div>
      </div>
      <div className="bg-white py-16 px-4">
        <div className="max-w-6xl mx-auto">
          <motion.h2
            className="text-3xl md:text-4xl font-bold text-center mb-12 text-black"
            initial={{ opacity: 0 }}
            whileInView={{ opacity: 1 }}
            transition={{ duration: 0.6 }}
            viewport={{ once: true }}
          >
            Why Choose ChatSphere?
          </motion.h2>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
            {features.map((feature, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.6, delay: index * 0.2 }}
                viewport={{ once: true }}
              >
                <Card className="text-black hover:bg-black/5 duration-200 ">
                  <CardHeader>
                    <div className="text-4xl mb-4 flex items-center justify-center">{feature.icon}</div>
                    <CardTitle className="text-xl font-semibold text-center">{feature.title}</CardTitle>
                  </CardHeader>
                  <CardContent>
                    <p className="text-gray-400 text-center">{feature.description}</p>
                  </CardContent>
                </Card>
              </motion.div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
}