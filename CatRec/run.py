#!/usr/bin/env python
# coding: utf-8

# In[5]:


import matplotlib.pyplot as plt
from PIL import Image
import os
import shutil
import tensorflow as tf
from tensorflow import keras
from tensorflow.keras import layers
from tensorflow.keras.models import Sequential
import numpy as np
from tensorflow.keras.models import load_model
import sys

# In[6]:


model = load_model('model.h5')


# In[7]:


listt=['https://baike.baidu.com/item/%E6%96%B0%E5%8A%A0%E5%9D%A1%E7%8C%AB/671184?fr=aladdin#2',
'https://baike.baidu.com/item/%E8%B1%B9%E7%8C%AB/4921581?fr=aladdin',
'https://baike.baidu.com/item/%E4%BC%AF%E6%9B%BC%E7%8C%AB/641852?fr=aladdin',
'https://baike.baidu.com/item/%E5%AD%9F%E4%B9%B0%E7%8C%AB/4510178?fr=aladdin',
'https://baike.baidu.com/item/%E8%8B%B1%E5%9B%BD%E7%9F%AD%E6%AF%9B%E7%8C%AB/672846?fr=aladdin',
'https://baike.baidu.com/item/%E5%9F%83%E5%8F%8A%E7%8C%AB/386503?fr=aladdin',
'https://baike.baidu.com/item/%E7%BC%85%E5%9B%A0%E7%8C%AB/647590?fr=aladdin',
'https://baike.baidu.com/item/%E6%B3%A2%E6%96%AF%E7%8C%AB/585?fr=aladdin',
'https://baike.baidu.com/item/%E9%9B%AA%E9%9E%8B%E7%8C%AB/4513092',
'https://baike.baidu.com/item/%E4%BF%84%E7%BD%97%E6%96%AF%E8%93%9D%E7%8C%AB/643065?fr=aladdin',
'https://baike.baidu.com/item/%E6%9A%B9%E7%BD%97%E7%8C%AB/556082?fr=aladdin',
'https://baike.baidu.com/item/%E5%8A%A0%E6%8B%BF%E5%A4%A7%E6%97%A0%E6%AF%9B%E7%8C%AB/643507?fromtitle=%E6%96%AF%E8%8A%AC%E5%85%8B%E6%96%AF%E7%8C%AB&fromid=8028531&fr=aladdin']


# In[12]:


batch_size = 32
img_height = 224
img_width = 224
data_dir = './cats'
train_ds = tf.keras.preprocessing.image_dataset_from_directory(
  data_dir,
  validation_split=0.2,
  subset="training",
  seed=123,
  image_size=(img_height, img_width),
  batch_size=batch_size)

val_ds = tf.keras.preprocessing.image_dataset_from_directory(
  data_dir,
  validation_split=0.2,
  subset="validation",
  seed=123,
  image_size=(img_height, img_width),
  batch_size=batch_size)


# In[14]:


class_names = train_ds.class_names


# In[15]:


img = Image.open(sys.argv[1])
img_array = keras.preprocessing.image.img_to_array(img)
img_array = tf.expand_dims(img_array, 0) 

predictions = model.predict(img_array)
score = tf.nn.softmax(predictions[0])

#print(
#    "This image most likely belongs to {} with a {:.2f} percent confidence."
#    .format(class_names[np.argmax(score)], 100 * np.max(score))
#)
print(class_names[np.argmax(score)])
#os.system('"C:\Windows\SystemApps\Microsoft.MicrosoftEdge_8wekyb3d8bbwe/MicrosoftEdge.exe" %s'%(listt[int(class_names[np.argmax(score)])]))


# In[ ]:




